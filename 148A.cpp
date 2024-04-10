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

	int k, l, m, n, d;
	
	cin >> k >> l >> m >> n;
	cin >> d;
	int ans = 0;
	for (int i = 1; i <= d; i++) {
		if (i % k == 0 || i % l == 0 || i % m == 0 || i % n == 0) {
			ans++;
		}
	}
	cout << ans << endl;

	return 0;
}