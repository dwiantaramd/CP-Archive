#include <bits/stdc++.h>

using namespace std;

int main() {
    long x ,y;
    cin >> x >> y;

    string str;
    cin >> str;

    for(int i = 0; i < str.length();i++) {
        if(str[i] == 't') {
            x -= 1;
        } else if(str[i]== 'b') {
            x += 1;
        } else if(str[i]== 'u') {
            y -= 1;
        } else if(str[i] == 's') {
            y += 1;
        }
    }

    cout << x << " " << y;
}
