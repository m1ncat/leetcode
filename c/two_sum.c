#include <stdio.h>
#include <stdlib.h>

int *twoSum(int *nums, int numsSize, int target)
{
	int i, j, index1, index2;

	for (i = 0; i < numsSize; ++i) {
		for (j = i + 1; j < numsSize; ++j) {
			if (*(nums + i) + *(nums + j) == target) {
				index1 = i;
				index2 = j;
				break;
			}
		}
	}

	int *res = malloc(sizeof(int) * 2);
	*res = index1;
	*(res + 1) = index2;

	return res;
}

int main()
{
	int nums[4] = {2, 7, 11, 15};
	int numsSize = sizeof(nums) / sizeof(int);
	int target = 9;
	int *res = twoSum(nums, numsSize, target);

	printf("res[0]: %d, res[1]: %d\n", res[0], res[1]);

	free(res);
	return 0;
}