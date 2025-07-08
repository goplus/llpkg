package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Sets the key used for AES encryption/decryption
 *
 * @param key pointer to the key
 * @param key_bytes number of bytes in key
 * @param mode key mode, 0 : decrypt, 1: encrypt
 *
 * @return uint8_t number of key bytes written to hardware, used for fault injection check
 */
//go:linkname AesHalSetkey C.aes_hal_setkey
func AesHalSetkey(key *c.Uint8T, key_bytes c.SizeT, mode c.Int) c.Uint8T

/**
 * @brief encrypts/decrypts a single block
 *
 * @param input_block input block, size of AES_BLOCK_BYTES
 * @param output_block output block, size of AES_BLOCK_BYTES
 */
//go:linkname AesHalTransformBlock C.aes_hal_transform_block
func AesHalTransformBlock(input_block c.Pointer, output_block c.Pointer)

/**
 * @brief Inits the AES mode of operation
 *
 * @param mode mode of operation, e.g. CTR or CBC
 */
// llgo:link EspAesModeT.AesHalModeInit C.aes_hal_mode_init
func (recv_ EspAesModeT) AesHalModeInit() {
}

/**
 * @brief Sets the initialization vector for the transform
 *
 * @note The same IV must never be reused with the same key
 *
 * @param iv the initialization vector, length = IV_BYTES (16 bytes)
 */
//go:linkname AesHalSetIv C.aes_hal_set_iv
func AesHalSetIv(iv *c.Uint8T)

/**
 * @brief Reads the initialization vector
 *
 * @param iv initialization vector read from HW, length = IV_BYTES (16 bytes)
 */
//go:linkname AesHalReadIv C.aes_hal_read_iv
func AesHalReadIv(iv *c.Uint8T)

/**
 * @brief Busy waits until the AES operation is done
 *
 * @param output pointer to inlink descriptor
 */
//go:linkname AesHalWaitDone C.aes_hal_wait_done
func AesHalWaitDone()

/**
 * @brief Starts an already configured AES DMA transform
 *
 * @param num_blocks Number of blocks to transform
 */
//go:linkname AesHalTransformDmaStart C.aes_hal_transform_dma_start
func AesHalTransformDmaStart(num_blocks c.SizeT)

/**
 * @brief Finish up a AES DMA conversion, release DMA
 *
 */
//go:linkname AesHalTransformDmaFinish C.aes_hal_transform_dma_finish
func AesHalTransformDmaFinish()
