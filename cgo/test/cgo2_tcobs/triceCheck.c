/*! \file triceCheck.c
\brief wrapper for tests
\author thomas.hoehenleitner [at] seerose.net
*******************************************************************************/
#include "trice.h"

// TriceCheck performs trice code sequence n.
// This function is called from Go for tests.
// Each trice line needs to have a commented // "string" with its expected output.
// The ID values must be in this file, because it is compiled first and trice update runs later.
void TriceCheck( int n ){
    switch( n ){
        default:
        break; case  0: TRICE8_1( Id(12222), "value=%d", -1 ); //exp: `value=-1`
        break; case  1: TRICE8_1( Id(11111), "value=%u", -1 ); //exp: `value=255`
    }
}
