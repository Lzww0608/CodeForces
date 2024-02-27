
#include <iostream>
#include <algorithm>
#include <cmath>
#include <sstream>
#include <string>
#include <vector>
#include <unordered_map>
using namespace std;



int main() {
	std::ios::sync_with_stdio(false);
	cin.tie(nullptr);
	int t, n;
	cin >> t;
	while (t--) {
		cin >> n;
		vector<int> pos(2 * n + 1);
		int ans = 0, x;
		for (int i = 1; i <= n; ++i) {
			cin >> x;
			pos[x] = i;
			for (int j = 1; j * x <= 2 * i; ++j) {
				if (j != x && pos[j] > 0 && pos[j] + i == j * x)
					ans++;
			}
		}

		cout << ans << endl;
	}
	return 0;
}


