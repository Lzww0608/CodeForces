#include <iostream>
#include <vector>
#include <unordered_map>
#include <cstdlib>
#include <cmath>
#include <algorithm>
#include <string>
using namespace std;
using ll = long long;


int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	cout.tie(nullptr);

	int k, r;
	cin >> k >> r;
	int ans = k;
	for (int i = 1; i <= 10; i++) {
		ans = k * i;
		if (ans % 10 == 0 || (ans - r) % 10 == 0) {
			break;
		}
	}

	cout << ans / k << endl;

	return 0;
}