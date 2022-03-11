<template>
  <section>
    <v-navigation-drawer v-model="drawer" app>
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title class="title">
            {{getUserActive}}
          </v-list-item-title>
          <v-list-item-subtitle>
            {{getRoleActive}}
          </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
      <v-divider></v-divider>
      <v-list dense nav>
        <v-list-item to="/dashboard">
          <v-list-item-action>
            <v-icon>mdi-home</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Dashboard</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item to="/">
          <v-list-item-action>
            <v-icon>mdi-account-group</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title>Pegawai</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item  to="/">
          <v-list-item-action>
            <v-icon>mdi-help-circle</v-icon>
          </v-list-item-action>
          <v-list-item-content>
                    <v-list-item-title>About</v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-list>
        <v-divider></v-divider>
        <v-list dense nav>
            <v-list-item  to="/Messages">
                <v-list-item-action>
                    <v-badge>
                        <template v-slot:badge>0</template>
                        <v-icon>mdi-email</v-icon>
                    </v-badge>
                </v-list-item-action>
                <v-list-item-content>
                    <v-list-item-title>Messages</v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-list>
        <template v-slot:append>
            <div class="pa-2">
                <v-btn @click="signoutButtonPressed" color="red darken-4 white--text" block>Logout</v-btn>
            </div>
        </template>
        </v-navigation-drawer>
 
        <v-app-bar app color="blue darken-4" dark>
            <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
            <v-toolbar-title>Personal History Card</v-toolbar-title>
            <div class="flex-grow-1"></div>
 
            <v-menu left bottom>
                <template v-slot:activator="{ on }">
                <v-btn icon v-on="on">
                    <v-icon>mdi-dots-vertical</v-icon>
                </v-btn>
                </template>
 
                <v-list>
                    <v-list-item  to="/Profile">
                        <v-list-item-action>
                            <v-icon>mdi-account-circle</v-icon>
                        </v-list-item-action>
                        <v-list-item-content>
                            <v-list-item-title>Profile</v-list-item-title>
                        </v-list-item-content>
                    </v-list-item>
                    <v-list-item  to="/Setting">
                        <v-list-item-action>
                            <v-icon>mdi-settings</v-icon>
                        </v-list-item-action>
                        <v-list-item-content>
                            <v-list-item-title>Setting</v-list-item-title>
                        </v-list-item-content>
                    </v-list-item>
                </v-list>
            </v-menu>
        </v-app-bar>
    </section>
</template>
<script>
  import { authenticationService } from '@/_services';
  import { Role } from '@/_helpers';
  export default {
      props: {
        source: String,
      },
      data: () => ({
        drawer: null,
        currentUser: null,
      }),
      methods: {
        signoutButtonPressed() { 
          authenticationService.logout();
          this.$router.push("/login");
        }
      },
      created () {
        authenticationService.currentUser.subscribe(x => this.currentUser = x);
      },
      computed: {
        isAdmin () {
            return this.currentUser && this.currentUser.level === Role.Admin;
        }, 
        getUserActive() {
          if (this.currentUser) {
            return this.currentUser.username;
          }
          return "Unknown"
        },
        getRoleActive() {
          if (this.currentUser) {
            return this.currentUser.level;
          }
          return "Unknown"
        }
      },
  };
</script>