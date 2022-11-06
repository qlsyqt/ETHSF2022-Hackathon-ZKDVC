import React, { useState } from "react";
import {
  Form,
  FormGroup,
  Modal,
} from "react-bootstrap";
import { toast } from "react-toastify";

export default function ClaimBadgeModal({ onHide, doClaim }) {
  // const [claim, setClaim] = useState("");
  const [aoxAddress, setAoxAddress] = useState("");

  const handleSubmit = async (event) => {
    console.log("onsubmit");
    event.preventDefault();
    if(!aoxAddress){
      toast.error('Address not valid');
      return
    }
    doClaim(aoxAddress)
    // replace with claimBadge function
    // const resp = await api.sendAuthProof(formData);
    // if (resp.code === 200) {
    //   setClaim(resp.result);
    //   toast.success("Claim download success!");
    // } else {
    //   toast.error("Failed:" + resp.message);
    // }
  };

  return (
    <Modal show={true} onHide={onHide} dialogClassName="w-64" centered={true}>
      <Modal.Header closeButton>
        <Modal.Title>Address</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form id="-area" onSubmit={handleSubmit}>
          <FormGroup controlId="addClaim">
            <Form.Control
              as="input"
              className="fs-6"
              value={aoxAddress}
              onChange={e => setAoxAddress(e.target.value)}
              name="input"
              required
            ></Form.Control>
          </FormGroup>
          <div className="text-right mt-4">
            <button
              className="text-white py-1 px-3 bg-[#f0504f] rounded "
              onClick={handleSubmit}
            >
              Done
            </button>
          </div>
        </Form>
      </Modal.Body>
    </Modal>
  );
}
