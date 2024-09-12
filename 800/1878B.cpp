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
		ll x = 1;
		for (int i = 1; i <= n; i++) {
			cout << x << " ";
			x += 2;
		}
		cout << '\n';
		
	}



	return 0;
}
