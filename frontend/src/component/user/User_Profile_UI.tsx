import * as React from 'react';
import { Button, Container } from "@mui/material";
import { Box, Grid, Paper, Link } from '@mui/material';
import Moment from 'moment';

import { useParams } from 'react-router-dom';

import { UsersInterface } from '../../models/user/IUser';
import ip_address from '../ip';

function User_Profile(){
    const { email } = useParams(); // ดึง parameter จาก url-parameter
    const [isDataLoaded, setIsDataloaded] = React.useState<boolean | null>(false);
    const [user, setUser] = React.useState<Partial<UsersInterface>>({});

    const [dialogUpdateOpen, setDialogUpdateOpen] = React.useState(false);

    Moment.locale('th');

    const handleDialogUpdateClickOpen = () => {
        setDialogUpdateOpen(true);
      };

    const getUser = async () => {
        const apiUrl = "http://" + ip_address() + ":8080/user/"+email; // email คือ email ที่ผ่านเข้ามาทาง parameter
        console.log(apiUrl)
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

    React.useEffect(() => {
        const fetchData = async () => {
            await getUser();
            setIsDataloaded(true);
        }
        fetchData();
    }, []);

    if(isDataLoaded) return (
        <Container>
            <Box>
                <Paper elevation={2} sx={{padding:2,margin:2}}>
                    <Grid container>
                        <Grid margin={2}>
                            <img src={`${user.Profile_Picture}`} width="250" height="250"/> {/** show base64 picture from string variable (that contain base64 picture data) */}
                        </Grid>
                        <Grid marginLeft={1} item xs={8}>
                            <Grid>
                                <h2>{user.Profile_Name}</h2>
                            </Grid>
                            <Grid item> {/** เอา Grid มาล็อคไม่ให้ component มันเด้งไปที่อื่น */}
                                <Box
                                    component="div"
                                    sx={{
                                    width : "100%",
                                    textOverflow: 'ellipsis',overflow: 'hidden',
                                    whiteSpace: 'break-spaces',
                                    my: 2,
                                    p: 1,
                                    bgcolor: (theme) =>
                                        theme.palette.mode === 'dark' ? '#101010' : 'grey.100',
                                    color: (theme) =>
                                        theme.palette.mode === 'dark' ? 'grey.300' : 'grey.800',
                                    border: '1px solid',
                                    borderColor: (theme) =>
                                        theme.palette.mode === 'dark' ? 'grey.800' : 'grey.300',
                                    borderRadius: 2,
                                    fontSize: '0.875rem',
                                    fontWeight: '700',
                                    }}
                                >{/** กำหนดให้เว้นบรรทัด auto จาก white space */}
                                    {"Full name: " + user.FirstName + " " + user.LastName + "\n\n"} 
                                    {"Email: " + user.Email+ "\n\n"}
                                    {"Birthday: " + `${Moment(user.Birthday).format('DD MMMM YYYY')}` + "\n\n"}
                                    {"Phone number: " + user.Phone_number}
                                </Box>
                            </Grid>
                            <Grid container>
                                <Grid item xs={6}>
                                    <Button
                                        type="submit"
                                        fullWidth
                                        variant="contained"
                                        sx={{ mt: 1, mb: 2 }}
                                        onClick={handleDialogUpdateClickOpen}
                                    >
                                        Update
                                    </Button>
                                </Grid>
                                <Grid item xs={6}>
                                    <Button
                                        type="submit"
                                        fullWidth
                                        variant="contained"
                                        sx={{ mt: 1, mb: 2 }}
                                        onClick={handleDialogUpdateClickOpen}
                                    >
                                        Change password
                                    </Button>
                                </Grid>
                            </Grid>
                        </Grid>

                    </Grid>
                </Paper>
            </Box>
        </Container>
    );
    return null;
}

export default User_Profile
