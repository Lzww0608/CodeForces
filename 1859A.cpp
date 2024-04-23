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
	

	ll n, t, x;
	cin >> t;
	while (t--) {
		cin >> n;
		vector<ll> a(n);
		ll mn = INT_MAX, cnt = 0;
		for (ll& x : a) {
			cin >> x;
			if (x < mn) {
				mn = x;
				cnt = 1;
			}
			else if (x == mn) {
				cnt++;
			}
		}
		if (cnt == n) {
			cout << -1 << endl;
			continue;
		}
		cout << cnt << ' ' << n - cnt << endl;
		for (int i = 0; i < cnt; i++) {
			cout << mn << ' ';
		}
		cout << endl;

		for (int x : a) {
			if (x != mn) {
				cout << x << ' ';
			}
		}
		cout << endl;
	}


	return 0;
}