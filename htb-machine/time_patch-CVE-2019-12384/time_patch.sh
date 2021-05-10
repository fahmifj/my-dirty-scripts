#!/bin/bash
project_path="/opt/json_project/"
new_jackson="jackson-core-2.11.0.jar"
old_jackson="jackson-core-2.9.8.jar"

if [ "$1" == "patch" ]; then
    
    cd /tmp
    wget "http://$2/$new_jackson"
    wget "http://$2/parse.rb"

    # backup the original code
    mv $project_path$old_jackson $project_path"classpath/"${old_jackson//.jar/.bak}
    mv $project_path"parse.rb" $project_path"parse.rb.bak"

    # move the updated parser and jackson
    mv /tmp/$new_jackson $project_path"classpath/"$new_jackson 
    mv /tmp/parse.rb $project_path"parse.rb.bak"

elif [ "$1" == "restore" ]; then

    rm /tmp/"$new_jackson"
    rm /tmp/"parse.rb"
    rm $project_path$new_jackson
    rm $project_path"parse.rb"

    mv $project_path"classpath/"${old_jackson//.jar/.bak} $project_path"classpath/"$old_jackson
    mv $project_path"parse.rb.bak" $project_path"parse.rb"
fi
