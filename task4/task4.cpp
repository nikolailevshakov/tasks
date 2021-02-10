#include <iostream>
#include <string>
#include <fstream>
#include <vector>
#include <algorithm>
#include <sstream>

class Person
{
public:
    int in;
    int out;
    int count;
    Person(int inn, int outt) :
        in(inn), out(outt), count(1) {}
    Person(int inn, int outt, int countt) :
        in(inn), out(outt), count(countt) {}
    
};


int time_to_mins(std::string str)
{
    int h, m;
    h = std::stoi(str.substr(0, str.find(":")));
    m = std::stoi(str.substr(str.find(":") + 1, str.size() - 1));
    return ((h - 8) * 60 + m);
}

std::string mins_to_time(int mins)
{
    int h = (mins / 60) + 8;
    int m = mins % 60;
    std::string time = std::to_string(h) + ":" + std::to_string(m);
    if (m == 0)
        time += "0";
    return time;
}


std::vector<Person> get_input(std::string input_path)
{
    std::vector<Person> people;
    std::ifstream input_file;
    input_file.open(input_path);
    if (!input_file) {
        std::cerr << "Can't open the input file" << std::endl;
        exit(1);
    }
    std::string t1, t2, line;
    while (std::getline(input_file, line))
    {
        std::istringstream ss(line);
        ss >> t1 >> t2;
        people.push_back(Person(time_to_mins(t1), time_to_mins(t2)));
    }
    return people;
}

void print_results(std::vector<Person> vect_c)
{
    for (Person p : vect_c)
    {
        std::cout << mins_to_time(p.in) << " " << mins_to_time(p.out) << std::endl;
    }
}


int main(int argc, char* argv[])
{
    argc = 1;
    std::string input = argv[1];
    std::vector<Person> vect;
    vect = get_input(input);

    //getting vector of the smallest time intervals
    bool repeat = true;
    bool flag = false;
    int count;
    while (repeat)
    {
        flag = false;
        count = 0;
        for (int i = 0; !flag && i < vect.size(); ++i)
        {
            for (int k = 0; !flag && k < vect.size(); ++k)
            {
                if (vect[i].in > vect[k].in && vect[i].in<vect[k].out && vect[i].out > vect[k].out)
                {
                    vect.push_back(Person(vect[k].in, vect[i].in));
                    vect.push_back(Person(vect[i].in, vect[k].out));
                    vect.push_back(Person(vect[i].in, vect[k].out));
                    vect.push_back(Person(vect[k].out, vect[i].out));
                    if (k > i)
                    {
                        vect.erase(vect.begin() + i);
                        vect.erase(vect.begin() + k - 1);
                    }
                    if (k < i)
                    {
                        vect.erase(vect.begin() + k);
                        vect.erase(vect.begin() + i - 1);
                    }
                    flag = true;
                    count++;
                }

                if (vect[i].in<vect[k].out && vect[i].in>vect[k].in && vect[i].out < vect[k].out)
                {
                    vect.push_back(Person(vect[k].in, vect[i].in));
                    vect.push_back(Person(vect[i].in, vect[i].out));
                    vect.push_back(Person(vect[i].in, vect[i].out));
                    vect.push_back(Person(vect[i].out, vect[k].out));
                    if (k > i)
                    {
                        vect.erase(vect.begin() + i);
                        vect.erase(vect.begin() + k - 1);
                    }
                    if (k < i)
                    {
                        vect.erase(vect.begin() + k);
                        vect.erase(vect.begin() + i - 1);
                    }
                    flag = true;
                    count++;
                }

                if (vect[i].out == vect[k].out && vect[i].in > vect[k].in && vect[i].in < vect[k].out)
                {
                    vect.push_back(Person(vect[k].in, vect[i].in));
                    vect.push_back(Person(vect[i].in, vect[k].out));
                    vect.push_back(Person(vect[i].in, vect[k].out));
                    if (k > i)
                    {
                        vect.erase(vect.begin() + i);
                        vect.erase(vect.begin() + k - 1);
                    }
                    if (k < i)
                    {
                        vect.erase(vect.begin() + k);
                        vect.erase(vect.begin() + i - 1);
                    }
                    flag = true;
                    count++;
                }

                if (vect[i].in == vect[k].in && vect[i].out > vect[k].in && vect[i].out < vect[k].out)
                {
                    vect.push_back(Person(vect[i].in, vect[i].out));
                    vect.push_back(Person(vect[i].in, vect[i].out));
                    vect.push_back(Person(vect[i].out, vect[k].out));
                    if (k > i)
                    {
                        vect.erase(vect.begin() + i);
                        vect.erase(vect.begin() + k - 1);
                    }
                    if (k < i)
                    {
                        vect.erase(vect.begin() + k);
                        vect.erase(vect.begin() + i - 1);
                    }
                    flag = true;
                    count++;
                }

            }
        }
        if (count == 0)
            repeat = false;
    }

    // counting repeated intervals
    flag = false;
    count = 1;
    while (count)
    {
        count = 0;
        flag = false;
        for (int i = 0; !flag && i < vect.size(); i++)
        {
            for (int k = 0; !flag && k < vect.size(); k++)
            {
                if (vect[i].in == vect[k].in && vect[i].out == vect[k].out)
                {
                    if (i != k)
                    {
                        vect[i].count += 1;
                        vect.erase(vect.begin() + k);
                        flag = true;
                        count++;
                    }

                }
            }
        }
    }

    // getting the most busy intervals in final_vect
    int max_count=1;
    for (Person p : vect)
    {
        if (p.count > max_count)
            max_count = p.count;
    }

    std::vector<Person> final_vect;

    for (Person p : vect)
    {
        if (p.count == max_count)
            final_vect.push_back(p);
    }

    // summing of adjacent intervals 
    count=1;
    flag=false;
    while (count)
    {
        count = 0;
        flag = false;
        for (int i = 0; !flag && i < final_vect.size(); ++i)
        {
            for (int k = 0; !flag && k < final_vect.size(); ++k)
            {
                if (final_vect[i].out == final_vect[k].in)
                {
                    final_vect.push_back(Person(final_vect[i].in, final_vect[k].out, final_vect[i].count));
                    if (k > i)
                    {
                        final_vect.erase(final_vect.begin() + i);
                        final_vect.erase(final_vect.begin() + k - 1);
                    }
                    if (k < i)
                    {
                        final_vect.erase(final_vect.begin() + k);
                        final_vect.erase(final_vect.begin() + i - 1);
                    }
                    flag = true;
                    count++;
                }
            }
        }
    }

    //grouping intervals in ascending order
    for (int i = 0; i < final_vect.size(); ++i)
    {
        for (int k = 0; k < final_vect.size(); ++k)
        {
            if (final_vect[i].in<final_vect[k].in && i>k)
            {
                std::swap(final_vect[i], final_vect[k]);
                break;
            }
        }
    }

    print_results(final_vect);
}



