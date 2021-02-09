#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <algorithm>
#include <sstream>

std::vector<double> get_input(std::string file_name, std::string input_pathl)
{
    std::ifstream input_file;
    input_file.open(input_pathl+"\\"+file_name);
    if (!input_file) {
        std::cerr << "Can't open the input file" << std::endl;
        exit(1);
    }
    double x;
    std::vector<double> means;
    std::string line;
    while (std::getline(input_file, line))
    {
        std::istringstream ss(line);
        ss >> x;
        means.push_back(x);
    }
    return means;
}

double get_max(std::vector<double> data)
{
    sort(data.begin(), data.end());
    return data[data.size() - 1];
}


int main(int argc, char* argv[])
{
    argc = 1;
    std::string input_path = argv[1];
    std::vector<double> cash1, cash2, cash3, cash4, cash5, sums;
    double sum = 0;

    cash1 = get_input("cash1.txt", input_path);
    cash2 = get_input("cash2.txt", input_path);
    cash3 = get_input("cash3.txt", input_path);
    cash4 = get_input("cash4.txt", input_path);
    cash5 = get_input("cash5.txt", input_path);

    for (int i = 0; i < cash1.size(); ++i)
    {
        sum = cash1[i] + cash2[i] + cash3[i] + cash4[i] + cash5[i];
        sums.push_back(sum);
        sum = 0;
    }

    double max_el = get_max(sums);
    for (int j = 0; j < sums.size(); ++j)
    {
        if (max_el == sums[j])
        {
            std::cout << j + 1 << std::endl;
        }
    }
    system("PAUSE");
}



