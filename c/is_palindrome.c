#include <stdio.h>
#include <stdbool.h>

bool isPalindrome(int x)
{
	int value = x;
    int length = 1;
    while (x / 10 > 0) {
    	length++;
    	x = x / 10;
    }

    char arr[length];
    sprintf(arr, "%d", value);

    int i, j;
    int firstPart[length / 2];
    int lastPart[length / 2];
    if (length % 2 == 0) {
    	for (i = 0; i < length / 2; ++i) {
    		firstPart[i] = arr[i];
    	}
    	int k = 0;
    	for (j = length - 1; j >= length / 2; j--) {
    		lastPart[k] = arr[j];
    		k++; 
    	}
    } else {
    	for (i = 0; i < length / 2; ++i) {
    		firstPart[i] = arr[i];
    	}
    	int k = 0;
    	for (j = length - 1; j > length / 2; j--) {
    		lastPart[k] = arr[j];
    		k++; 
    	}
    }

    bool flag = false;
    printf("len: %d\n", length / 2);
    for (i = 0; i < length / 2; i++) {
    	printf("firstPart: %d, lastPart: %d\n", firstPart[i], lastPart[i]);
    	if (firstPart[i] != lastPart[i]) {
    		flag = false;
    		break;
    	} else {
    		flag = true;
    	}
    }

    if (length == 1) {
    	flag = true;
    }

    if (value < 0) {
    	flag = false;
    }

    return flag;
}

int main() 
{
	int test = 1000021;
	bool flag = isPalindrome(test);
	printf("%s\n", flag?"true":"false");
}