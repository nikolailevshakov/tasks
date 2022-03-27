#include <iostream>
#include <vector>
#include <fstream>
#include <string>
#include <algorithm>
#include <iomanip>

void printd(double a)
{
    std::cout << std::fixed;
    std::cout << std::setprecision(2);
    std::cout << a << std::endl;
}

std::vector<int> get_input(std::string input_way)
{
    std::ifstream input;
    input.open(input_way);
    if (!input) {
        std::cerr << "Can't open the input file" << std::endl;
        system("PAUSE");
        exit(1);
    }
    int x;
    std::vector<int> digits;
    while (input >> x)
    {
        digits.push_back(x);
    }
    return digits;
}

double get_mean(std::vector<int> data)
{
    double sum = 0, mean;
    for (int digit : data)
    {
        sum += digit;
    }
    return sum / data.size();
}

double get_min(std::vector<int> data)
{
    std::sort(data.begin(), data.end());
    return data[0];
}

double get_max(std::vector<int> data)
{
    std::sort(data.begin(), data.end());
    return data[data.size() - 1];
}

double get_median(std::vector<int> data)
{
    std::sort(data.begin(), data.end());
    double median;
    if (data.size() % 2 == 0)
        return (data[data.size() / 2] + data[data.size() / 2 - 1]) / 2;
    if (data.size() % 2 == 1)
        return data[(int)(data.size() / 2)];
}

double get_perc90(std::vector<int> data)
{
    std::sort(data.begin(), data.end());
    double perc;
    int x;
    if (data.size() == 10)
    {
        return (double)(data[9] - data[8]) / 10 + data[8];
    }
    x = (int)data.size() * 0.9;
    perc = data[x-1];
    return perc;

}


int main(int argc, char* argv[])
{
    argc = 1;
    std::string input_way = argv[1];
    std::vector<int> input_data = get_input(input_way);

    printd(get_perc90(input_data));
    printd(get_median(input_data));
    printd(get_max(input_data));
    printd(get_min(input_data));
    printd(get_mean(input_data));

    system("PAUSE");
}



