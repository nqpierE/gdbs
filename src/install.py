import os 

print("[*] building...")
os.chdir("src")
os.system("go build .")
os.system("mv gdbs ../")
os.chdir("../")
print("[*] done.")
print("[-] run following command")
root_dir = os.getcwd()
print("[-] $ echo 'export PATH=$PATH:" + root_dir + "' >> ~/.zshrc")
