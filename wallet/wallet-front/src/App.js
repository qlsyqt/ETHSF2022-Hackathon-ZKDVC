import { useEffect, useState } from "react";
import * as api from "./api/api";
import "./App.css";
import { scanUrl } from "./config";
import { sleep } from "./utils";
import IconBgShadow from "./assets/bg-shadow.png";
import IconKnn3 from "./assets/knn3.png";
import IconAdd from "./assets/add.png";
import ClaimBadgeModal from "./components/ClaimBadgeModal";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import {
  Form,
  FormGroup,
  Modal,
} from "react-bootstrap";

const SUCCESS = 200;

function App() {
  const [did, setDid] = useState("");
  const [authModalVisible, setAuthModalVisible] = useState(false);
  const [claimModalVisible, setClaimModalVisible] = useState(false);
  const [activeClaimItem, setActiveClaimItem] = useState({});

  const onClaim = (item) => {
    setActiveClaimItem(item);
    setClaimModalVisible(true);
  };

  const doClaim = (aoxAddress) => {
    console.log("aox address is: ", aoxAddress);
    api.acquirerBadge(activeClaimItem.hIndex).then(async (resp) => {
      await sleep(1000);
      processResponse(resp, () => {
        window.location.reload();
      });
    });
  };

  useEffect(() => {
    document.title = "Demo Wallet";
    api.fetchPolygonId().then((resp) => {
      processResponse(resp, () => setDid(resp.result));
    });
  }, []);

  return (
    <div className="App bg-[#f8f9fa]">
      <header className="App-header py-8 shadow-go bg-white">
        <div className="container flex gap-4 items-center">
          <div className="font-bold text-2xl lg:text-3xl">Wallet</div>
        </div>
      </header>
      <div className="relative pt-10 pb-24 min-h-screen">
        <img
          src={IconBgShadow}
          className="absolute w-16 h-16 md:w-48 md:h-48 right-0 top-0 transform rotate-180"
        />
        <img
          src={IconBgShadow}
          className="absolute w-16 h-16 md:w-48 md:h-48 left-0 bottom-0"
        />
        <div className="container">
          <div className="mb-8">
            {/* <h1 className="text-xl lg:text-2xl font-bold mb-6">Polygon Wallet</h1> */}
            <WalletArea did={did} />
          </div>
          <div className="mb-8">
            <div className="flex items-center justify-between mb-6">
              <h1 className="text-xl lg:text-2xl font-bold">Your Claims</h1>
              <img
                src={IconAdd}
                className="w-14 h-14 cursor-pointer"
                onClick={() => setAuthModalVisible(true)}
              />
            </div>
            <ClaimsArea onClaim={onClaim} />
          </div>
          <div className="mb-8">
            <h1 className="text-xl lg:text-2xl font-bold mb-6">Your Badges</h1>
            <BadgesArea />
          </div>
        </div>
      </div>

      <ToastContainer position="top-center" autoClose={2000} />
      {authModalVisible && (
        <AuthJsonArea onHide={() => setAuthModalVisible(false)} />
      )}
      {claimModalVisible && (
        <ClaimBadgeModal
          onHide={() => setClaimModalVisible(false)}
          doClaim={doClaim}
        />
      )}
    </div>
  );
}

const DataCategoryBadge = ({ text }) => (
  <>
    {text === "0" ? (
      <div className="py-0.5 px-2 rounded inline-block text-white bg-[#518EFF]">
        ENS
      </div>
    ) : text === "1" ? (
      <div className="py-0.5 px-2 rounded inline-block text-white bg-[#056AE7]">
        NFT
      </div>
    ) : text === "2" ? (
      <div className="py-0.5 px-2 rounded inline-block text-white bg-[#FFC100]">
        Snapshot
      </div>
    ) : text === "3" ? (
      <div className="py-0.5 px-2 rounded inline-block text-[#00511E] bg-[#ABFE2C]">
        LENS
      </div>
    ) : (
      <div className="py-0.5 px-2 rounded inline-block text-white bg-black">
        {text}
      </div>
    )}
  </>
);

function WalletArea({ did }) {
  const [wallet, setWallet] = useState({
    mainWalletAddress: "loading",
    badgeWalletAddress: "loading",
    network: "loading",
  });

  useEffect(() => {
    api.fetchPolygonWalletInfo().then((resp) => {
      processResponse(resp, () => {
        setWallet(resp.result);
      });
    });
  }, []);

  return (
    <div className="bg-white shadow-go rounded-2xl px-6 py-2">
      <div className="flex items-center my-3">
        <div className="w-64">Address</div>
        <div className="font-bold text-[#9f9f9f]">
          {wallet.mainWalletAddress}
        </div>
      </div>
      {/* <div className="flex items-center my-3">
        <div className="w-64">Aux Wallet Address</div>
        <div className="font-bold text-[#9f9f9f]">
          {wallet.auxWalletAddress}
        </div>
      </div> */}
      <div className="flex items-center my-3">
        <div className="w-64">ID</div>
        <div className="font-bold text-[#9f9f9f]">{did}</div>
      </div>
      <div className="flex items-center my-3">
        <div className="w-64">Network</div>
        <div className="font-bold text-[#9f9f9f]">{wallet.network}</div>
      </div>
    </div>
  );
}

function ClaimsArea({ onClaim }) {
  const [claims, setClaims] = useState([]);

  useEffect(() => {
    api.fetchClaims().then((resp) =>
      processResponse(resp, () => {
        setClaims([...resp.result]);
      })
    );
  }, []);

  // var AddArea = showAdd?(<AuthJsonArea></AuthJsonArea>):(<div></div>);
  const ClaimItem = ({ item }) => (
    <div className="w-full md:w-1/2 lg:w-1/4 p-3 lg:p-6 lg:-p-6">
      <div className="shadow-go rounded-2xl bg-white overflow-hidden">
        <div className="flex py-2 px-4 items-center justify-between border-b border-[#a5a5a5] bg-red-gradient">
          <div className="text-2xl font-bold longtext">#{item.hIndex}</div>
          <div>
            <DataCategoryBadge text={item.dataCategory} />
          </div>
        </div>
        <div className="pt-2 px-3 pb-3">
          <div className="flex items-center gap-2 my-2">
            <div className="text-[#6d757e] font-bold">
              {item.dataCategory === "1" && <>Contract</>}
              {item.dataCategory === "2" && <>Space ID</>}:{" "}
            </div>
            {item.dataCategory === "1" ? (
              <a
                href={`${scanUrl}/address/${item.subCategory}`}
                target="_blank"
                className="bg-[#f9f9f9] text-black no-underline rounded px-2 longtext"
              >
                {item.subCategory}
              </a>
            ) : (
              <div className="bg-[#f9f9f9] text-black  rounded px-2">
                {item.subCategory}
              </div>
            )}
          </div>
          <div className="flex items-center gap-2 my-2">
            <div className="text-[#6d757e] font-bold">Interval:</div>
            <div>{item.interval}</div>
          </div>
          <div className="text-right mt-2">
            <button
              className="text-white py-1 px-3 bg-[#f0504f] rounded "
              onClick={() => onClaim(item)}
            >
              Claim Badge
            </button>
          </div>
        </div>
      </div>
    </div>
  );

  return (
    <div>
      <div className="flex flex-wrap -m-3 flex-col md:flex-row">
        {claims.map((item) => (
          <ClaimItem item={item} />
        ))}
      </div>
    </div>
  );
}

function AuthJsonArea({ onHide }) {
  var [claim, setClaim] = useState("");
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async (event) => {
    if (submitting) {
      return;
    }
    try {
      setSubmitting(true);
      console.log("onsubmit");
      event.preventDefault();
      const inputData = document.getElementById("authjson").value;
      const resp = await api.sendAuthProof(inputData);
      setSubmitting(false);
      if (resp.code === 200) {
        setClaim(resp.result);
        toast.success("Claim download success!");
        window.location.reload(); //自动刷新
      } else {
        toast.error("Failed:" + resp.message);
      }
    } catch (err) {
      setSubmitting(false);
    }
  };

  return (
    <Modal show={true} onHide={onHide} dialogClassName="w-64" centered={true}>
      <Modal.Header closeButton>
        <Modal.Title>Select Auth Json</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form id="auth-area" onSubmit={handleSubmit}>
          <FormGroup controlId="addClaim">
            <Form.Control
              id="authjson"
              as="textarea"
              rows={5}
              className="fs-6"
              name="input"
              required
            ></Form.Control>
          </FormGroup>
          <div className="text-right mt-4">
            <button
              disabled={submitting}
              className="text-white py-1 px-3 bg-[#f0504f] rounded "
              onClick={handleSubmit}
            >
              {submitting ? "Submitting" : "Submit"}
            </button>
          </div>
        </Form>
      </Modal.Body>
    </Modal>
  );
}

function BadgesArea() {
  const [badges, setBadges] = useState([]);
  useEffect(() => {
    api.fetchBadges().then((resp) =>
      processResponse(resp, () => {
        setBadges(resp.result);
      })
    );
  }, []);

  const BadgeItem = ({ item }) => (
    <div className="w-full md:w-1/2 lg:w-1/4 p-3">
      <div className="shadow-go rounded-2xl bg-white">
        <div className="text-center py-7">
          <img src={IconKnn3} className="w-32 mx-auto mb-2" />
          <DataCategoryBadge text={item.dataCategory} />
        </div>
        <div className="px-3 py-2 border-t border-[#979797]">
          <div className="flex items-center gap-2 my-2">
            <div className="text-[#6d757e] font-bold">Token ID:</div>
            <div className="longtext">{item.badge}</div>
          </div>
          <div className="flex items-center gap-2 my-2">
            <div className="text-[#6d757e] font-bold">Badge Contract:</div>
            <div className="longtext">{item.badgeContract}</div>
          </div>
        </div>
      </div>
    </div>
  );

  return (
    <div>
      <div className="flex flex-wrap flex-col -m-3 md:flex-row">
        {badges.map((item) => (
          <BadgeItem key={item.badge} item={item} />
        ))}
      </div>
    </div>
  );
}

function processResponse(resp, callback) {
  if (resp.code === SUCCESS) {
    callback();
  } else {
    toast.error(resp.message);
  }
}

export default App;
