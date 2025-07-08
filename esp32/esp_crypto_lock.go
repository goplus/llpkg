package freertos

import _ "unsafe"

/**
 * @brief Acquire lock for HMAC cryptography peripheral
 *
 * Internally also locks the SHA peripheral, as the HMAC depends on the SHA peripheral
 */
//go:linkname EspCryptoHmacLockAcquire C.esp_crypto_hmac_lock_acquire
func EspCryptoHmacLockAcquire()

/**
 * @brief Release lock for HMAC cryptography peripheral
 *
 * Internally also releases the SHA peripheral, as the HMAC depends on the SHA peripheral
 */
//go:linkname EspCryptoHmacLockRelease C.esp_crypto_hmac_lock_release
func EspCryptoHmacLockRelease()

/**
 * @brief Acquire lock for DS cryptography peripheral
 *
 * Internally also locks the HMAC (which locks SHA), AES and MPI  peripheral, as the DS depends on these peripherals
 */
//go:linkname EspCryptoDsLockAcquire C.esp_crypto_ds_lock_acquire
func EspCryptoDsLockAcquire()

/**
 * @brief Release lock for DS cryptography peripheral
 *
 * Internally also releases the HMAC (which locks SHA), AES and MPI  peripheral, as the DS depends on these peripherals
 */
//go:linkname EspCryptoDsLockRelease C.esp_crypto_ds_lock_release
func EspCryptoDsLockRelease()

/**
 * @brief Acquire lock for the SHA and AES cryptography peripheral.
 *
 */
//go:linkname EspCryptoShaAesLockAcquire C.esp_crypto_sha_aes_lock_acquire
func EspCryptoShaAesLockAcquire()

/**
 * @brief Release lock for the SHA and AES cryptography peripheral.
 *
 */
//go:linkname EspCryptoShaAesLockRelease C.esp_crypto_sha_aes_lock_release
func EspCryptoShaAesLockRelease()

/**
 * @brief Acquire lock for the mpi cryptography peripheral.
 *
 */
//go:linkname EspCryptoMpiLockAcquire C.esp_crypto_mpi_lock_acquire
func EspCryptoMpiLockAcquire()

/**
 * @brief Release lock for the mpi/rsa cryptography peripheral.
 *
 */
//go:linkname EspCryptoMpiLockRelease C.esp_crypto_mpi_lock_release
func EspCryptoMpiLockRelease()
