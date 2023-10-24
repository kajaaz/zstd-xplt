#include <stdlib.h>
#include "zstd/zstd_compress.c"

size_t ZSTD_getSequences(ZSTD_CCtx* zc, ZSTD_Sequence* outSeqs, size_t outSeqsSize, const void* src, size_t srcSize);
