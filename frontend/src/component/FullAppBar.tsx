import * as React from 'react';
import AppBar from '@mui/material/AppBar';
import Box from '@mui/material/Box';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import AccountCircle from '@mui/icons-material/AccountCircle';
import MenuItem from '@mui/material/MenuItem';
import Menu from '@mui/material/Menu';
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import Drawer from '@mui/material/Drawer';
import TwitterIcon from '@mui/icons-material/Twitter';

// User Icon
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import HomeIcon from '@mui/icons-material/Home';

// Admin Icon
import ViewCarouselIcon from '@mui/icons-material/ViewCarousel';
import TrendingUpTwoToneIcon from '@mui/icons-material/TrendingUpTwoTone';

import { Link as RouterLink } from "react-router-dom";
import { UsersInterface } from '../models/user/IUser';
// import { AdminsInterface } from '../models/admin/IAdmin';

function FullAppBar() {
  const [user, setUser] = React.useState<Partial<UsersInterface>>({});
  // const [admin, setAdmin] = React.useState<Partial<AdminsInterface>>({});
  const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };

  function drawerList() {
    if (localStorage.getItem("position") == "Admin") {
    } else { // User Drawer
      return (
        <List sx={{ width: "100%" }}>

          <ListItem button component={RouterLink} to="/">
            <HomeIcon />
            <ListItemText primary="FirstPage" sx={{ paddingLeft: 1 }} />
          </ListItem>

        </List>
      );
    }
  }

  // function myProfileUser() {
  //   if (localStorage.getItem("position") == "User")
  //     return (
  //       <MenuItem onClick={handleClose} component={RouterLink} to={"/user_profile/" + localStorage.getItem("id")} >My Profile</MenuItem>
  //     )
  // }

  const [auth] = React.useState(true);
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);

  const handleMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const [isDrawerOpen, setIsDrawerOpen] = React.useState(false);

  const getUser = async () => {
    const apiUrl = "http://192.168.1.37:8080/user/" + localStorage.getItem("email");
    const requestOptions = {
      method: "GET",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
    };

    await fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          setUser(res.data);

        }
      });
  };

  // async function GetAdmin() {
  //   const apiUrl = "http://localhost:8080/admin/" + localStorage.getItem("email");
  //   const requestOptions = {
  //     method: "GET",
  //     headers: {
  //       Authorization: `Bearer ${localStorage.getItem("token")}`,
  //       "Content-Type": "application/json",
  //     },
  //   };

  //   await fetch(apiUrl, requestOptions)
  //     .then((response) => response.json())
  //     .then((res) => {
  //       if (res.data) {
  //         setAdmin(res.data);
  //       }
  //     });
  // };

  React.useEffect(() => {
    const fetchData = async () => {
      await getUser();
      // await GetAdmin();
    }
    fetchData();
  }, []);

  return (
    <Box sx={{ flexGrow: 1 }}>

      <AppBar position="static">
        <Toolbar>
          <IconButton
            edge="start"
            color="inherit"
            aria-label="menu"
            onClick={() => setIsDrawerOpen(true)}
          >
            <MenuIcon />
          </IconButton>

          <Drawer open={isDrawerOpen} onClose={() => setIsDrawerOpen(false)}>

            <TwitterIcon color="primary" sx={{ fontSize: 150, margin: 1, padding: 2 }} />
            {/** List of Drawer Divided by position */}
            {drawerList()}

          </Drawer>

          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Entacc
          </Typography>

          {auth && (                                                                               /* รูป Icon Profild */
            <div>
              <IconButton
                size="large"
                aria-label="account of current user"
                aria-controls="menu-appbar"
                aria-haspopup="true"
                onClick={handleMenu}
                color="inherit"
              >
                <AccountCircle />
              </IconButton>
              <Menu
                id="menu-appbar"
                anchorEl={anchorEl}
                anchorOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                keepMounted
                transformOrigin={{
                  vertical: 'top',
                  horizontal: 'right',
                }}
                open={Boolean(anchorEl)}
                onClose={handleClose}
              >
                {/* {myProfileUser()} for user */}
                <MenuItem onClick={signout} component={RouterLink} to="/" >Logout</MenuItem>
              </Menu>
            </div>
          )}

        </Toolbar>
      </AppBar>

    </Box>

  );
}

export default FullAppBar;