#include <cstdio>
#include <cstring>
#include <vector>
#include <iostream>
#include <cmath>
#include <algorithm>
using namespace std;
typedef long long ll;

int main(){
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr); 
	cout.tie(nullptr);

	int q;
	cin >> q;
	while (q--) {
		int n, k;
		cin >> n >> k;
		vector<int> a(k), t(k);
		vector<ll> f(n, 1e18);
		for (int& x : a) {
			cin >> x;
			x--;
		}
		for (int i = 0; i < k; ++i) {
			cin >> t[i];
			f[a[i]] = t[i];
		}

		for (int i = n - 2; i >= 0; i--) {
			f[i] = min(f[i], f[i + 1] + 1);
		}

		ll mn = 1e18;
		for (ll x : f) {
			mn = min(mn + 1, x);
			cout << mn << " ";
		}
		cout << '\n';
	}
	return 0;
}