package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Initialize GDMA channel retention link for powerdown the TOP powerdomain during lightsleep
 * @param  group_id Group id
 * @param  pair_id  Pair id
 * @return
 *      - ESP_OK: Create DMA retention link successfully
 *      - ESP_ERR_NO_MEM: Create DMA retention link failed because out of memory
 */
//go:linkname GdmaSleepRetentionInit C.gdma_sleep_retention_init
func GdmaSleepRetentionInit(group_id c.Int, pair_id c.Int) EspErrT

/**
 * Destroy GDMA channel retention link
 * @param  group_id Group id
 * @param  pair_id  Pair id
 * @return
 *      - ESP_OK: GDMA channel retention link destrory successfully
 *      - ESP_ERR_INVALID_STATE: GDMA channel retention link not create yet
 */
//go:linkname GdmaSleepRetentionDeinit C.gdma_sleep_retention_deinit
func GdmaSleepRetentionDeinit(group_id c.Int, pair_id c.Int) EspErrT
