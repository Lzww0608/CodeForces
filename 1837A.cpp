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

const int mod = 1e9 + 7;

int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	cout.tie(nullptr);
	

	ll t, k, x;
	cin >> t;
	while (t--) {
		cin >> x >> k;
		if (x % k != 0) {
			cout << 1 << endl;
			cout << x << endl;
			continue;
		}
		cout << 2 << endl;
		cout << 1 << ' ' << x - 1 << endl;
	}


	return 0;
}