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
	
	ll n, x;
	cin >> n >> x;
	ll ans = x, cnt = 0;
	char c;
	for (int i = 0; i < n; i++) {
		cin >> c;
		cin >> x;
		if (c == '+') {
			ans += x;
		}
		else if (ans >= x) {
			ans -= x;
		}
		else {
			cnt++;
		}
	}

	cout << ans << ' ' << cnt << endl;


	return 0;
}
