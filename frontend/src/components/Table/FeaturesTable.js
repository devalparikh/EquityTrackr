import React from "react";

import {
  TableHeader,
  TableCell,
  TableContainer,
  Input,
} from "@windmill/react-ui";

import SectionTitle from "../Typography/SectionTitle";

function FeaturesTable({
  checkedFeatures,
  features,
  title,
  updateCheckedFeatures,
}) {
  return (
    <>
      <SectionTitle>{title}</SectionTitle>
      <TableContainer className="mb-8 p-4">
        <TableHeader>
          {features.map((feature, index) => (
            <label className="inline-flex items-center">
              <Input
                type="checkbox"
                checked={checkedFeatures[feature]}
                onChange={() => updateCheckedFeatures(feature, index)}
              />
              <TableCell
                className={`pl-1 ${
                  !checkedFeatures[feature] && "text-gray-400"
                }`}
                key={index}
              >
                {feature}
              </TableCell>
            </label>
          ))}
        </TableHeader>
      </TableContainer>
    </>
  );
}
export default FeaturesTable;
