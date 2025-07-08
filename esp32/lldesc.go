package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const LLDESC_TX_MBLK_SIZE = 268
const LLDESC_RX_SMBLK_SIZE = 64
const LLDESC_RX_MBLK_SIZE = 524
const LLDESC_RX_AMPDU_ENTRY_MBLK_SIZE = 64
const LLDESC_RX_AMPDU_LEN_MBLK_SIZE = 256
const LLDESC_TX_MBLK_NUM = 10
const LLDESC_RX_MBLK_NUM = 10
const LLDESC_RX_AMPDU_ENTRY_MBLK_NUM = 4
const LLDESC_RX_AMPDU_LEN_MLBK_NUM = 8
const LLDESC_OWNER_MASK = 0x80000000
const LLDESC_OWNER_SHIFT = 31
const LLDESC_SW_OWNED = 0
const LLDESC_HW_OWNED = 1
const LLDESC_EOF_MASK = 0x40000000
const LLDESC_EOF_SHIFT = 30
const LLDESC_SOSF_MASK = 0x20000000
const LLDESC_SOSF_SHIFT = 29
const LLDESC_LENGTH_MASK = 0x00fff000
const LLDESC_LENGTH_SHIFT = 12
const LLDESC_SIZE_MASK = 0x00000fff
const LLDESC_SIZE_SHIFT = 0
const LLDESC_ADDR_MASK = 0x000fffff

/**
 * Generate a linked list pointing to a (huge) buffer in an descriptor array.
 *
 * The caller should ensure there is enough size to hold the array, by calling
 * ``lldesc_get_required_num_constrained`` with the same max_desc_size argument.
 *
 * @param[out] out_desc_array Output of a descriptor array, the head should be fed to the DMA.
 * @param buffer Buffer for the descriptors to point to.
 * @param size Size (or length for TX) of the buffer
 * @param max_desc_size Maximum length of each descriptor
 * @param isrx The RX DMA may require the buffer to be word-aligned, set to true for a RX link, otherwise false.
 */
// llgo:link (*LldescT).LldescSetupLinkConstrained C.lldesc_setup_link_constrained
func (recv_ *LldescT) LldescSetupLinkConstrained(buffer c.Pointer, size c.Int, max_desc_size c.Int, isrx bool) {
}

/**
 * @brief Get the received length of a linked list, until end of the link or eof.
 *
 * @param head      The head of the linked list.
 * @param[out] out_next Output of the next descriptor of the EOF descriptor. Return NULL if there's no
 *                 EOF. Can be set to NULL if next descriptor is not needed.
 * @return The accumulation of the `len` field of all descriptors until EOF or the end of the link.
 */
// llgo:link (*LldescT).LldescGetReceivedLen C.lldesc_get_received_len
func (recv_ *LldescT) LldescGetReceivedLen(out_next **LldescT) c.Int {
	return 0
}

type TxAmpduEntryS struct {
	SubLen  c.Uint32T
	DiliNum c.Uint32T
	c.Uint32T
	NullByte c.Uint32T
	Data     c.Uint32T
	Enc      c.Uint32T
	Seq      c.Uint32T
}
type TxAmpduEntryT TxAmpduEntryS

type LldescChainS struct {
	Head *LldescT
	Tail *LldescT
}
type LldescChainT LldescChainS
