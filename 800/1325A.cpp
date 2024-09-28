#include <iostream>
#include <vector>
#include <unordered_map>
#include <cstdlib>
#include <cmath>
#include <algorithm>
#include <string>
using namespace std;
using ll = long long;

int lcm(int x, int y) {
	int res = x * y;
	while (y > 0) {
		int tmp = x;
		x = y;
		y = tmp % y;
	}
	return res / x;
}

int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	cout.tie(nullptr);
	
	int t, x;
	cin >> t;
	while (t--) {
		cin >> x;
		cout << 1 << ' ' << x - 1 << endl;

	}


	return 0;
}
