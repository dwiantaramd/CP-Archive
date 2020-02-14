
#include <iostream>
#include <string>
#define ll long long
using namespace std;

bool isalphabet(char a){
    return ((a >= 'a' && a <= 'z') ||(a >= 'A' && a <= 'Z'));
}
int main(){
    ll n;
    cin>>n;
    cin.ignore();
    for(ll i = 0; i < n; i++){
        string s;
        getline(cin,s);
        ll panjang = s.length();
        ll mid = (panjang)/2;
        if(panjang % 2 == 0)
            mid--;
        string solve = "";
        for(ll j = 0; j < panjang; j++){
            if(j <= mid)
                if(isalphabet(s[j]))
                    solve = (char)((int) s[j] + 2) + solve;
                else
                    solve = (char)((int)s[j] - 1) + solve;
            else
                if(isalphabet(s[j]))
                    solve = (char)((int) s[j] + 3) + solve;
                else
                    solve = s[j] + solve;
        }
        cout<<solve<<endl;
    }

}
