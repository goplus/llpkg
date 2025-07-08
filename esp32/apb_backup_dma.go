package freertos

import _ "unsafe"

/*
 * SPDX-FileCopyrightText: 2019-2021 Espressif Systems (Shanghai) CO LTD
 *
 * SPDX-License-Identifier: Apache-2.0
 */
//go:linkname EtsApbBackupInitLockFunc C.ets_apb_backup_init_lock_func
func EtsApbBackupInitLockFunc(_apb_backup_lock func(), _apb_backup_unlock func())
