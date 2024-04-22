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
		ll mn = INT_MAX, cnt = 0;
		for (int i = 0; i < n; i++) {
			cin >> x;
			if (x < mn) {
				cnt = 1;
				mn = x;
			}
			else if (x == mn) {
				cnt++;
			}
		}

		cout << n - cnt << endl;
	}


	return 0;
}