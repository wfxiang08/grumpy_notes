global_meta_b = {}
def hello_b(name, metas):
    print "hello, world: %s %s" % (name, metas)


def get_key_value_b(reqs, key):
    return "result: %s %s" % (reqs.get(key, ""), global_meta_b)
