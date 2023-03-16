import * as React from 'react';
import { DataGrid, GridColDef, GridValueGetterParams } from '@mui/x-data-grid';
import { Grid, Paper, Table, TableCell, TableContainer, TableHead, TableRow } from '@mui/material';

import ip_address from '../ip';
import { AccountsInterface } from '../../models/account/IAccount';

export default function All_Account_UI() {
    const [account, setAccount] = React.useState<AccountsInterface[]>([]);

    const columns: GridColDef[] = [
        { field: 'ID_Account', headerName: 'ID', width: 70},
        { field: 'Twitter_Account', headerName: 'Twitter Account', width: 200 },
        { field: 'Twitter_Password', headerName: 'Twitter password', width: 200 },
        { field: 'Email', headerName: 'Email', width: 200 },
        { field: 'Email_Password', headerName: 'Email password', width: 200 },
        { field: 'Phone_Number', headerName: 'Phone number', width: 130 },
        { field: 'Years', headerName: 'Years', width: 90, },
        { field: 'Account_Status', headerName: 'Status', width: 130, valueGetter: (params) => params.value.Status, },
      ];   

    const getAccount = async () => {
        const apiUrl = "http://" + ip_address() + ":8080/account/"+localStorage.getItem('email'); // email คือ email ที่ผ่านเข้ามาทาง parameter
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
                    setAccount(res.data); 
                }
            });
    };

    React.useEffect(() => {
        const fetchData = async () => {
            await getAccount();
            await console.log(account[0].Account_Status)
        }
        fetchData();
    }, []);

    return (
        <Grid>
            <div style={{ height: 540, width: '100%' }}>
                <DataGrid
                    rows={account}
                    getRowId={(row) => row.ID_Account}
                    columns={columns}
                />
            </div>
        </Grid>
    );
}