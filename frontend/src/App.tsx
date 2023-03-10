
import './App.css';
import SignIn from './components/SignIn';
import MuiAppBar, { AppBarProps as MuiAppBarProps } from "@mui/material/AppBar";
import { styled, createTheme, ThemeProvider } from "@mui/material/styles";
import MuiDrawer from "@mui/material/Drawer";
import HomeIcon from "@mui/icons-material/Home";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import React, { useState, useEffect } from "react";
import Box from "@mui/material/Box";
import CssBaseline from "@mui/material/CssBaseline";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import Divider from "@mui/material/Divider";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import Container from "@mui/material/Container";
import Home from './components/Home';
import MenuBookRoundedIcon from '@mui/icons-material/MenuBookRounded';
import Patient from './components/Patient/PatientList';
import PatientCreate from './components/Patient/PatientCreate';
import PatientEdit from './components/Patient/PatientEdit';
import Prescription from './components/Prescription/PrescriptionList';
import PrescriptionDelete from './components/Prescription/PrescriptionDelete';
import PrescriptionCreate from './components/Prescription/PrescriptionCreate';
import PrescriptionEdit from './components/Prescription/PrescriptionEdit';
import BoyIcon from '@mui/icons-material/Boy';
import ReceiptLongIcon from '@mui/icons-material/ReceiptLong';

const drawerWidth = 240;

interface AppBarProps extends MuiAppBarProps {
  open?: boolean;
}

const AppBar = styled(MuiAppBar, {
  shouldForwardProp: (prop) => prop !== "open",
})<AppBarProps>(({ theme, open }) => ({
  zIndex: theme.zIndex.drawer + 1,
  transition: theme.transitions.create(["width", "margin"], {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  ...(open && {
    marginLeft: drawerWidth,
    width: `calc(100% - ${drawerWidth}px)`,
    transition: theme.transitions.create(["width", "margin"], {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
  }),
}));

const Drawer = styled(MuiDrawer, {
  shouldForwardProp: (prop) => prop !== "open",
})(({ theme, open }) => ({
  "& .MuiDrawer-paper": {
    position: "relative",
    whiteSpace: "nowrap",
    width: drawerWidth,
    transition: theme.transitions.create("width", {
      easing: theme.transitions.easing.sharp,
      duration: theme.transitions.duration.enteringScreen,
    }),
    boxSizing: "border-box",
    ...(!open && {
      overflowX: "hidden",
      transition: theme.transitions.create("width", {
        easing: theme.transitions.easing.sharp,
        duration: theme.transitions.duration.leavingScreen,
      }),
      width: theme.spacing(7),
      [theme.breakpoints.up("sm")]: {
        width: theme.spacing(9),
      },
    }),
  },
}));

const mdTheme = createTheme();

const menu = [
  { name: "?????????????????????", icon: <HomeIcon />, path: "/", role: "Nurse"}, 
  { name: "????????????????????????????????????????????????", icon: <BoyIcon />, path: "/patients/list" ,role: "Nurse"},
  //{ name: "????????????????????????????????????????????????", icon: <MenuBookRoundedIcon />, path: "/books",role: "admin"},
  { name: "????????????????????????", icon: <ReceiptLongIcon/>, path: "/prescription/list",role: "Nurse"},
  //{ name: "??????????????????????????????????????????", icon: <MeetingRoomIcon />, path: "/researchroomreservationrecords" ,role: "user"},
  //{ name: "???????????????????????????????????????????????????????????????", icon: <ComputerIcon />, path: "/computer_reservations" ,role: "user"},
  //{ name: "??????????????????????????????", icon: <AddCircleIcon />, path: "/borrows",role: "admin" },
  //{ name: "????????????????????????????????????????????????????????????", icon: <StorefrontIcon />, path: "/bills" ,role : "admin"},
  //{ name: "Problem Report", icon: <FlagIcon />, path: "/problemreports" ,role: "user"},
  // { name: "?????????????????????", icon: <PeopleIcon />, path: "/books",role:"employee"},
  // { name: "??????????????????", icon: <PeopleIcon />, path: "/users",role:"student"},
  //{ name: "?????????????????????????????????????????????", icon: <YouTubeIcon />, path: "/watch_videos" },
];
function App() {
  const [token, setToken] = useState<String>("");
  const [open, setOpen] = React.useState(true);
  //const [email, setEmail] = useState<string | null>();
  const [role, setRole] = useState<String| null>();
  const toggleDrawer = () => {
    setOpen(!open);
  };

  useEffect(() => {
    const token = localStorage.getItem("token");
    //const email= localStorage.getItem("email");
    const role = localStorage.getItem("role")
    //console.log(email)
    if (token) {
      setToken(token);
      setRole(role);
    }
  }, []);
  

  if (!token) {
    return <SignIn />;
  }

  const signout = () => {
    localStorage.clear();
    window.location.href = "/";
  };
  
  return (
    <Router>
      <ThemeProvider theme={mdTheme}>
        <Box sx={{ display: "flex" }}>
          <CssBaseline />
          <AppBar position="absolute" open={open}>
            <Toolbar
              sx={{
                pr: "24px", // keep right padding when drawer closed
              }}
            >
              <IconButton
                edge="start"
                color="inherit"
                aria-label="open drawer"
                onClick={toggleDrawer}
                sx={{
                  marginRight: "36px",
                  ...(open && { display: "none" }),
                }}
              >
                <MenuIcon />
              </IconButton>
              <Typography
                component="h1"
                variant="h6"
                color="inherit"
                noWrap
                sx={{ flexGrow: 1 }}
              >
                ?????????????????????????????????????????????????????????
              </Typography>
              <Button color="inherit" onClick={signout}>
                ??????????????????????????????
              </Button>
            </Toolbar>
          </AppBar>
          <Drawer variant="permanent" open={open}>
            <Toolbar
              sx={{
                display: "flex",
                alignItems: "center",
                justifyContent: "flex-end",
                px: [1],
              }}
            >
              <IconButton onClick={toggleDrawer}>
                <ChevronLeftIcon />
              </IconButton>
            </Toolbar>
            <Divider />
            <List>
              {menu.map((item, index) => (
                 //email == item.email &&
                 role == item.role &&
                <Link
                  to={item.path}
                  key={item.name}
                  style={{ textDecoration: "none", color: "inherit" }}
                >
                  <ListItem button>
                    <ListItemIcon>{item.icon}</ListItemIcon>
                    <ListItemText primary={item.name} />
                  </ListItem>
                </Link>

              ))}
            </List>
          </Drawer>
          <Box
            component="main"
            sx={{
              backgroundColor: (theme) =>
                theme.palette.mode === "light"
                  ? theme.palette.grey[100]
                  : theme.palette.grey[900],
              flexGrow: 1,
              height: "100vh",
              overflow: "auto",
            }}
          >
            <Toolbar />
            <Container maxWidth="lg" sx={{ mt: 4, mb: 4 }}>
              <Routes>
                <Route path="/" element={<Home/>} />
                <Route path="/patients/list" element={<Patient/>} />
                <Route path="/patient/create" element={<PatientCreate/>} />
                <Route path="/patient/edit" element={<PatientEdit/>} />
                <Route path="/prescription/list" element={<Prescription/>} />
                <Route path="/prescription/delete" element={<PrescriptionDelete/>} />
                <Route path="/prescription/create" element={<PrescriptionCreate/>} />
                <Route path="/prescription/edit" element={<PrescriptionEdit/>} />
              </Routes>
            </Container>
          </Box>
        </Box>
      </ThemeProvider>
    </Router>
  );
}

export default App;
