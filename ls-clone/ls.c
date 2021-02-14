#include <sys/types.h>
#include <sys/stat.h>
#include <dirent.h>
#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <sys/types.h>
#include <pwd.h>
#include <grp.h>
#include <sys/syslimits.h>
#include <string.h>
#include <unistd.h>
#include <time.h>


#define CYAN "\x1b[36m" 
#define BLUE "\x1b[34m"
#define OFF "\x1b[0m"

char *getEntName(DIR *directory, struct dirent *ent) {
    char *entName;

    if (ent->d_type == DT_LNK) {
        char symLinkPath[PATH_MAX];
        size_t sizeLink = readlinkat(dirfd(directory), ent->d_name,symLinkPath,PATH_MAX);
        size_t sizeEntName = strlen(CYAN) + strlen(ent->d_name) + 4 + strlen(OFF) + sizeLink + 1;
        entName = malloc(sizeEntName);
        snprintf(entName, sizeEntName, CYAN"%s"OFF" -> %s", ent->d_name, symLinkPath);
    } else if (ent->d_type == DT_DIR) {
        size_t sizeEntName = strlen(BLUE) + strlen(ent->d_name) + strlen(OFF) + 1;
        entName = malloc(sizeEntName);
        snprintf(entName, sizeEntName, BLUE"%s"OFF, ent->d_name);
    } else {
        size_t sizeEntName = strlen(ent->d_name) + 1;
        entName = malloc(sizeEntName);
        snprintf(entName, sizeEntName, "%s", ent->d_name);
    }

    return entName;
}

void listDir(char *dirName) {
    DIR *directory = opendir(dirName);

    struct dirent **ents;
    struct passwd *user;
    struct group *group;
    struct stat stats;

    int numEnts = scandir(dirName, &ents, NULL, alphasort);

    for(int i = 0; i < numEnts; i++) {
        fstatat(dirfd(directory), ents[i]->d_name, &stats, 0);

        user = getpwuid(stats.st_uid);
        group = getgrgid(stats.st_gid);

        struct timespec ts = stats.st_mtimespec;
        struct tm *modified = localtime(&ts.tv_sec);
        char timeStr[17];
        strftime(timeStr, 16, "%Y-%m-%dT%H:%M", modified);

        off_t entSize = stats.st_size;

        char *entName = getEntName(directory, ents[i]);
        
        char *outputStr = "%8s %8s %10lld %10s %-128s\n\0";
        printf(outputStr,user->pw_name,group->gr_name, entSize, timeStr, entName);

        free(entName);
        free(ents[i]);
    }

    free(ents);
}

int main(int argc, char **argv) {
    char *dirName = ".";
    if (argc > 1) {
        dirName = argv[1];
    }

    listDir(dirName);
}
