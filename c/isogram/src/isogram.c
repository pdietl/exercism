#include <stddef.h>
#include <ctype.h>
#include <stdbool.h>
#include "isogram.h"

bool is_isogram(const char phrase[])
{
    size_t i;
    char ch;
    char counts[26] = {0};

    if (!phrase)
        return false;

    for (ch = *phrase; ch; ch = *++phrase) {
        if (isalpha(ch)) {
            counts[tolower(ch) - 'a']++;
        }
    }

    for (i = 0; i < sizeof counts; ++i) {
        if (counts[i] > 1)
            return false;
    }
    return true;
}
