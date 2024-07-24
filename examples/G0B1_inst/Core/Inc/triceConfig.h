/*! \file triceConfig.h
\author Thomas.Hoehenleitner [at] seerose.net
*******************************************************************************/

#ifndef TRICE_CONFIG_H_
#define TRICE_CONFIG_H_

#ifdef __cplusplus
extern "C" {
#endif

// hardware specific trice lib settings
#include "main.h" 
#define TriceStamp16  0x1616 // TIM17->CNT     // 0...999 us
#define TriceStamp32  0x32323232 // HAL_GetTick()  // 0...2^32-1 ms

//#define TRICE_BUFFER TRICE_RING_BUFFER

#define TRICE_SINGLE_MAX_SIZE 16
#define TRICE_DATA_OFFSET 16
#define TRICE_DEFERRED_BUFFER_SIZE 256
//#define TRICE_DEFERRED_XTEA_ENCRYPT 1
#define TRICE_DEFERRED_OUT_FRAMING TRICE_FRAMING_COBS

// trice l -p JLINK -args="-Device STM32G0B1RE -if SWD -Speed 4000 -RTTChannel 0" -pf none -ts ms -d16
//#define TRICE_DIRECT_OUTPUT 1
//#define TRICE_DIRECT_SEGGER_RTT_32BIT_WRITE 1

// trice log -p com1
#define TRICE_DEFERRED_OUTPUT 1
#define TRICE_DEFERRED_UARTA 1
#define TRICE_UARTA USART2

#include "cmsis_gcc.h"
#define TRICE_ENTER_CRITICAL_SECTION { uint32_t primaskstate = __get_PRIMASK(); __disable_irq(); {
#define TRICE_LEAVE_CRITICAL_SECTION } __set_PRIMASK(primaskstate); }

#ifdef __cplusplus
}
#endif

#endif /* TRICE_CONFIG_H_ */
