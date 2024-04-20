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
		ll h1 = n / 3, h2 = n / 3, h3 = n / 3;
		if (n % 3 == 0) {
			h1++;
			h3--;
		}
		else if (n % 3 == 1) {
			h1 += 2;
			h3--;
		}
		else {
			h1 += 2;
			h2++;
			h3--;
		}

		cout << h2 << ' ' << h1 << ' '<< h3 << endl;
	}



	return 0;
}