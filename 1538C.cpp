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

	int T, l, r, n;
	string s;
	cin >> T;
	while (T--) {
		ll ans = 0;
		cin >> n >> l >> r;
		vector<int> a(n);
		for (int &x : a) {
			cin >> x;
		}
		sort(a.begin(), a.end());
		for (int i = 0; i < n; ++i) {
			auto x = upper_bound(a.begin(), a.begin() + i, r - a[i]);
			auto y = lower_bound(a.begin(), a.begin() + i, l - a[i]);
			ans += x - y;
		}
		cout << ans << endl;
	}
	return 0;
}