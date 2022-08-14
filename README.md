## ANSIBLE PLAYBOOK FOR MARIADB MAXSCALE KEEPALIVED

### Requirements

    - Debian 11 "Bullseye"
    - Ansible playbook:
        ```yml
            ansible-playbook [core 2.13.1]
                config file = None
                configured module search path = ['/home/auzty/.ansible/plugins/modules', '/usr/share/ansible/plugins/modules']
                ansible python module location = /home/auzty/.local/lib/python3.8/site-packages/ansible
                ansible collection location = /home/auzty/.ansible/collections:/usr/share/ansible/collections
                executable location = /home/auzty/.local/bin/ansible-playbook
                python version = 3.8.10 (default, Mar 15 2022, 12:22:08) [GCC 9.4.0]
                jinja version = 3.1.2
                libyaml = True
        ```

## HOW To Run

 - Check the Hostlist on `__DEPLOY__`

 the format of `[mariadbhost]` are `nodename=maria1 master=yes` ( for master node , no need specify `master` on slave node)

 the format of `[maxscalehost]` are `maxscalenode=maxscale1 noderole=master` ( for the master node, and for setup keepalived)

 check on file `site.yml` , and set your password for mysql root password : `current_password: "passwordku"` and maxscale password on `maxscale_password: "maxscalepwd"`

 and then run the ansible playbook

 ```bash
    ansible-playbook -vv  -i __DEPLOY__ site.yml
 ```

 ## The Infrastructure:

  ```
  Mariadb 1 : 192.168.56.120
  Mariadb 2 : 192.168.56.121
  Mariadb 3 : 192.168.56.122

  Maxscale 1 : 192.168.56.130
  Maxscale 2 : 192.168.56.131

  Keepalived vInterface : 192.168.56.150
  ```