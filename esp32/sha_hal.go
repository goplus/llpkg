package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Hashes a single message block
 *
 * @param sha_type          SHA algorithm to hash with
 * @param data_block        Input message to be hashed
 * @param block_word_len    Length of the input message
 * @param first_block       Is this the first block in a message or a continuation?
 */
// llgo:link EspShaType.ShaHalHashBlock C.sha_hal_hash_block
func (recv_ EspShaType) ShaHalHashBlock(data_block c.Pointer, block_word_len c.SizeT, first_block bool) {
}

/**
 * @brief Polls and waits until the SHA engine is idle
 *
 */
//go:linkname ShaHalWaitIdle C.sha_hal_wait_idle
func ShaHalWaitIdle()

/**
 * @brief Reads the current message digest from the SHA engine
 *
 * @param sha_type SHA algorithm used
 * @param digest_state Output buffer to which to read message digest to
 */
// llgo:link EspShaType.ShaHalReadDigest C.sha_hal_read_digest
func (recv_ EspShaType) ShaHalReadDigest(digest_state c.Pointer) {
}

/**
 * @brief Writes the message digest to the SHA engine
 *
 * @param sha_type The SHA algorithm type
 * @param digest_state Message digest to be written to SHA engine
 */
// llgo:link EspShaType.ShaHalWriteDigest C.sha_hal_write_digest
func (recv_ EspShaType) ShaHalWriteDigest(digest_state c.Pointer) {
}

/**
 * @brief Hashes a number of message blocks using DMA
 *
 * @param sha_type      SHA algorithm to hash with
 * @param num_blocks    Number of blocks to hash
 * @param first_block   Is this the first block in a message or a continuation?
 */
// llgo:link EspShaType.ShaHalHashDma C.sha_hal_hash_dma
func (recv_ EspShaType) ShaHalHashDma(num_blocks c.SizeT, first_block bool) {
}

/**
 * @brief Calculates and sets the initial digiest for SHA512_t
 *
 * @param t_string
 * @param t_len
 */
//go:linkname ShaHalSha512InitHash C.sha_hal_sha512_init_hash
func ShaHalSha512InitHash(t_string c.Uint32T, t_len c.Uint8T)
