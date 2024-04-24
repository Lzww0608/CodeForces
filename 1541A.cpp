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
		if (n & 1) {
			cout << 3 << ' ' << 1 << ' ' << 2 << ' ';
			for (int i = 4; i <= n; i++) {
				if (i & 1) {
					cout << i - 1 << ' ';
				}
				else {
					cout << i + 1 << ' ';
				}
			}
		}
		else {
			for (int i = 1; i <= n; i++) {
				if (i & 1) {
					cout << i + 1 << ' ';
				}
				else {
					cout << i - 1 << ' ';
				}
			}
		}
		cout << endl;
	}


	return 0;
}