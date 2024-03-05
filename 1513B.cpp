#include <iostream>
#include <unordered_map>
#include <vector>
#include <utility>
#include <map>
using ll = long long;
using namespace std;

const ll mod = 1e9 + 7;

vector<ll> f(200005, 0);

ll factorial(ll x) {
	if (x == 0 || x == 1)
		return f[x] = 1;
	if (f[x] > 0)
		return f[x];
	return f[x] = x * factorial(x - 1) % mod;
}

int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int t;
	cin >> t;
	while (t--) {
		int n;
		cin >> n;
		vector<int> a(n);
		int x = -1;
		for (int i = 0; i < n; ++i) {
			cin >> a[i];
			x &= a[i];
		}
		ll cnt = 0;
		for (int num : a) {
			if (num == x)
				++cnt;
		}
		cout << cnt * (cnt - 1) % mod * factorial(n - 2) % mod << endl;
		/*ll ans = cnt * (cnt - 1) % mod;
		for (int i = 2; i <= n - 2; i++) {
			ans = ans * i % mod;
		}
		cout << ans << endl;*/
	}

	return 0;

}