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

	int T, a, b;
	string s;
	cin >> T;
	while (T--) {
		cin >> a >> b;
		cin >> s;
		int ans = 0;
		int n = s.length();
		for (int i = 0, last = -1; i < n; i++) {
			if (s[i] == '1') {
				if (~last) {
					ans += min(a, (i - last - 1) * b);
				} else {
					ans += a;
				}
				last = i;
			}
		}
		cout << ans << endl;
	}
	return 0;
}