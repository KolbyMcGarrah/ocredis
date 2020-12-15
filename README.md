# ocredis
Instruments gopkg.in/redis.v3 and gopkg.in/redis.v4 interactions with Open Census

# contributions
Other redis versions can be added by adding a folder with the new version, copying the wrapper.go file from a previous version into the new directory, importing the new redis version, and then making any updates to the calls if they've been changed. 

New calls can be added to older versions by updating the wrapper.go files in the respective versions and then adding any missing structs to the commands.go file. Any version specific commands can be added to the version that requires them
