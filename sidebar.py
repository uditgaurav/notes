import json

with open("version-1.9.0-sidebars.json",'r') as json_file:
    data = json.load(json_file)

data['version-1.8.0-docs'] = data.pop('docs')
docs = data['version-1.8.0-docs']

# print (docs['Getting Started'])

val= [] 
# update version in getting started
for values in docs['Getting Started']:
    val.append("version-1.8.0-"+values)
docs['Getting Started']=val

val= [] 
# update version in litmus demo
for values in docs['Litmus Demo']:
    val.append("version-1.8.0-"+values)
docs['Litmus Demo'] =val

val= [] 
# update version in concepts
for values in docs['Concepts']:
    val.append("version-1.8.0-"+values)
docs['Concepts'] =val

val= [] 
# update version in platform ids
for values in docs['Platforms']:
    for v in values['ids']:
        val.append("version-1.8.0-"+v)
        values['ids']=val
    val = []
    
val= [] 
# update version in experiment ids
for values in docs['Experiments']:
    for v in values['ids']:
        val.append("version-1.8.0-"+v)
        values['ids']=val
    val = []

val= [] 
# update version in scheduler
for values in docs['Scheduler']:
    val.append("version-1.8.0-"+values)
docs['Scheduler'] =val

val= [] 
# update version in chaos workflow
for values in docs['Chaos Workflow']:
    val.append("version-1.8.0-"+values)
docs['Chaos Workflow'] =val

val= [] 
# update version in litmus FAQs
for values in docs['Litmus FAQs']:
    val.append("version-1.8.0-"+values)
docs['Litmus FAQs'] =val

val= [] 
# update version in litmus FAQs
for values in docs['Advanced']:
    val.append("version-1.8.0-"+values)
docs['Advanced'] =val

with open("version-1.9.0-sidebars.json",'w') as json_file:
     json.dump(data,json_file,indent=2)
