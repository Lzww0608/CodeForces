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
	
	int t, n, m;
	cin >> t;
	while (t--) {
		cin >> n >> m;
		ll ans = 0;
		if (n == 2) {
			ans = m;
		}
		else if (n > 2) {
			ans = m * 2;
		}
		cout << ans << endl;
	}
	

	return 0;
}