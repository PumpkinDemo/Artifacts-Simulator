import requests
import sys
import json

# host = 'http://39.103.229.106:8080'
host = 'http://127.0.0.1:8080'


embryo = {}

def show_a_artifact(artifact):
    res = {}
    res['套装'] = artifact['套装']
    res['位置'] = artifact['位置']
    res['Name'] = artifact['Name']
    res['等级'] = artifact['Lv']
    
    res[artifact['MainStat']['词条']] = artifact['MainStat']['数值']
    res['副词条'] = {}
    for substat in artifact['SubStat']:
        if not substat['词条']:
            continue
        res['副词条'][substat['词条']] = substat['数值']
    print(json.dumps(res, indent=4, ensure_ascii=False))
    

def gain():
    url = host + '/gain'
    data = {
        "domain": "Peak of Vindagnyr", 
        "resin": 20
    }
    r = requests.post(url, json=data)
    res = r.json()[0]
    # res.pop('')
    show_a_artifact(res)
    globals()['embryo'] = res


def enhance():
    url = host + '/enhance'
    dogfoods = [
        {"stars":5, "lv": 20},
        {"stars":5, "lv": 20}
    ]
    if not embryo:
        print('no embryo set')
        return
    data = {
        "embryo": embryo,
        "dogfoods": dogfoods
    }
    r = requests.post(url, json=data)
    res = r.json()
    show_a_artifact(res)


if __name__ == '__main__':
    argc = len(sys.argv)
    argv = sys.argv
    last_action = '1'
    while 1:
        prompt = '\nartifact action:'
        prompt += '\n\t1. gain' + (' (default)' if last_action == '1' else '')
        prompt += '\n\t2. enhance' + (' (default)\n' if last_action == '2' else '\n')
        action = input(prompt)
        if action == '':
            action = last_action
        if action == '1':
            gain()
        elif action == '2':
            enhance()
        elif action == '-1':
            break
        last_action = action