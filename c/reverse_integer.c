#include <stdio.h>
#include <math.h>

int reverse(int x)
{
	int operateNo = 0;
	if (x >= 0) {
		operateNo = x;
	} else {
		operateNo = -x;
	}

	int quotient = operateNo / 10;
	int reminder = operateNo % 10;
	int arr[10] = {0, 0, 0, 0, 0, 0, 0, 0, 0, 0};

	int count = 0;
	while (quotient > 0) {
		arr[count] = reminder;
		reminder = quotient % 10;
		quotient = quotient / 10;
		count++;
	}

	arr[count] = reminder;

	int figure = 0;
	int i;

	for (i = 0; i < 10; ++i) {
		printf("arr: %d\n", arr[i]);
	}

	int first = 0;
	int last = 0;
	for (i = 0; i < 10; ++i) {
		if (arr[i] != 0) {
			first = i;
			break;
		}
	}

	for (i = 0; i < 10; ++i) {
		if (arr[i] != 0) {
			last = i;
		}
	}

	figure = last - first;

	int res = 0;

	for (i = first; i <= last; i++) {
		res += arr[i] * pow(10, figure);
		figure -= 1;
	}

	if (res < 0) {
		return 0;
	}

	int resfinal = 0;
	if (x >= 0) {
		resfinal = res;
	} else {
		resfinal = -res;
	}

	if (resfinal > pow(2, 32) - 1 || resfinal < -(pow(2, 32))) {
		return 0;
	} else {
		return resfinal;
	}
}

int main() 
{
	int test = 1534236469;
	int res = reverse(test);

	printf("%d\n", res);
	return 0;
}