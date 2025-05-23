// Copyright (c) 2021 Jorge Luis Betancourt. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.
//
// +build darwin,cgo

package sensor

/*
#cgo CFLAGS: -x objective-c -fmodules -fblocks
#cgo LDFLAGS: -framework CoreFoundation -framework LocalAuthentication -framework Foundation
#include <stdlib.h>
#include <stdio.h>
#import <LocalAuthentication/LocalAuthentication.h>

int isTouchIDAvailable() {
    int result = 0;
    bool success = [[[LAContext alloc] init] canEvaluatePolicy:LAPolicyDeviceOwnerAuthentication error:nil];
    if (success) {
        return 1;
    }

    return 0;
}

int Authenticate(char const* reason) {
  LAContext *myContext = [[LAContext alloc] init];
  NSError *authError = nil;
  dispatch_semaphore_t sema = dispatch_semaphore_create(0);
  NSString *nsReason = [NSString stringWithUTF8String:reason];
  __block int result = 0;

  if ([myContext canEvaluatePolicy:LAPolicyDeviceOwnerAuthentication error:&authError]) {
    [myContext evaluatePolicy:LAPolicyDeviceOwnerAuthentication
      localizedReason:nsReason
      reply:^(BOOL success, NSError *error) {
        if (success) {
          result = 1;
        } else {
          result = 2;
        }
        dispatch_semaphore_signal(sema);
      }];
  }

  dispatch_semaphore_wait(sema, DISPATCH_TIME_FOREVER);
  dispatch_release(sema);
  return result;
}
*/
import (
	"C"
)
import (
	"errors"
	"unsafe"
)

// IsTouchIDAvailable checks if Touch ID is available in the current device
func IsTouchIDAvailable() bool {
	result := C.isTouchIDAvailable()

	return result == 1
}

func Authenticate(reason string) (bool, error) {
	reasonStr := C.CString(reason)
	defer C.free(unsafe.Pointer(reasonStr))

	result := C.Authenticate(reasonStr)
	switch result {
	case 1:
		return true, nil
	case 2:
		return false, nil
	}

	return false, errors.New("Error occurred accessing biometrics")
}