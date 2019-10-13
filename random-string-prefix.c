#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

#define Length 6

const char * sample(const char set[]) {
    return &set[arc4random_uniform(strlen(set))];
}

int main(void) {
    const char letters[] = "bcdfghjlmnpqrstvwz";
    const char digits[] = "269";

    const bool noLeadingDigit = true;

    char all[strlen(letters) + strlen(digits)];

    snprintf(&all[0], sizeof(letters), "%s", letters);
    snprintf(&all[strlen(letters)], sizeof(digits), "%s", digits);

    const size_t len = Length;
    char out[len + 1];
    size_t i = 0;

    if (noLeadingDigit) {
        snprintf(&out[0], sizeof(char) * 2, "%c", *sample(letters));
        i++;
    }

    while (i < len ) {
        snprintf(&out[i++], sizeof(char) * 2, "%c", *sample(all));
    }

    printf("%s\n", out);
    return 0;
}
