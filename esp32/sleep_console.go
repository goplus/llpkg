package freertos

import _ "unsafe"

type SleepConsoleUsjEnableStateT struct {
	UsjClockEnabled bool
	UsjPadEnabled   bool
}

/**
 * @brief Disable usb-serial-jtag pad during light sleep to avoid current leakage and
 *        backup the enable state before light sleep
 */
//go:linkname SleepConsoleUsjPadBackupAndDisable C.sleep_console_usj_pad_backup_and_disable
func SleepConsoleUsjPadBackupAndDisable()

/**
 * @brief Restore initial usb-serial-jtag pad enable state when wakeup from light sleep
 */
//go:linkname SleepConsoleUsjPadRestore C.sleep_console_usj_pad_restore
func SleepConsoleUsjPadRestore()
