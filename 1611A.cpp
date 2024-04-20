#include <iostream>
#include <vector>
#include <unordered_map>
#include <cstdlib>
#include <cmath>
#include <queue>
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

const int mod = 1e9 + 7;

int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	cout.tie(nullptr);
	

	ll n, t;
	cin >> t;
	while (t--) {
		cin >> n;
		if ((n & 1) == 0) {
			cout << 0 << endl;
			continue;
		}
		bool f = false;
		while (n >= 10) {
			n /= 10;
			if ((n & 1) == 0) {
				f = true;
			}
		}
		if ((n & 1) == 0) {
			cout << 1 << endl;
		} else if (f) {
			cout << 2 << endl;
		}
		else {
			cout << -1 << endl;
		}
		
	}



	return 0;
}