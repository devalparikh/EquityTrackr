import React from "react";

import {
  TableBody,
  TableFooter,
  TableContainer,
  Table,
  TableHeader,
  TableCell,
  TableRow,
  Avatar,
  Badge,
} from "@windmill/react-ui";

import SectionTitle from "../Typography/SectionTitle";
import { currencyFormat } from "../../utils/textFormat";

function PositionsTable({ data }) {
  return (
    <>
      <SectionTitle>My Positions (x{data.length})</SectionTitle>
      <TableContainer>
        <Table>
          <TableHeader>
            <tr>
              <TableCell>Name</TableCell>
              <TableCell className="text-right">Market Value</TableCell>
            </tr>
          </TableHeader>
          <TableBody>
            {data.map((position, i) => (
              <TableRow key={i}>
                <TableCell>
                  <div className="flex items-center text-sm">
                    <Avatar
                      className="hidden mr-3 md:block"
                      src={position.avatar}
                      alt="position image"
                    />
                    <div>
                      <p className="font-semibold">{position.name}</p>
                      <p className="text-xs text-gray-600 dark:text-gray-400">
                        {position.job}
                      </p>
                    </div>
                  </div>
                </TableCell>
                <TableCell>
                  <div className="flex flex-col text-right text-sm">
                    <p className="font-semibold">
                      {currencyFormat(position.amount)}
                    </p>
                    <p className="text-xs">{currencyFormat(position.openPL)}</p>
                  </div>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
export default PositionsTable;
