#include <bits/stdc++.h>

using namespace std;

int main() {
    string x;
    getline(cin,x);

    long long temp;
    cin >> temp;
    temp = temp % 26;
    char kapt[26] = {'A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'};
    char alfa[26] = {'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'};
    for(size_t i = 0; i< x.length(); i++) {
        bool found1 = false;
        bool found2 = false;
        long long f1, f2 = 0;
        for(int j = 0; j < 26; j++) {
            if(x[i] == kapt[j]) {
                found1 = true;
                f1 = j + temp;
                long long temp1 = temp;
                if(f1> 25) {
                    temp1 = temp1 - (25-j);
                    f1 = -1 + temp1;
                } else if (f1 < 0) {
                    temp1 = temp1 + j;
                    f1 = 26 + temp1;
                }
            }
            if(x[i] == alfa[j]) {
                found2 = true;
                f2 = j+temp;
                long long temp2 = temp;
                if(f2> 25) {
                    temp2 = temp2 -(25-j);
                    f2 = -1 + temp2;
                } else if (f2 < 0) {
                    temp2 = temp2 + j;
                    f2 = 26 + temp2;
                }
            }
        }

        if(found1) {
            cout << kapt[f1];
        } else if (found2) {
            cout << alfa[f2];
        } else {
            cout << x[i];
        }
    }
}
