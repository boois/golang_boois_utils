import re
from collections import OrderedDict
args_str='''&s.word.0=123123&s.word1.1=asdfasd&&s.word1=asdfasd'''
args_list = re.findall("[&]+s\.([^=\.]*)(?:(\.0|\.1)*)?[^=]*\s*=\s*([^&]*)", args_str)
search_group = OrderedDict()
for word, is_show, val in args_list:
    search_group[word] = (word.strip(), is_show == ".1" or is_show == "", val)
print search_group.values()
