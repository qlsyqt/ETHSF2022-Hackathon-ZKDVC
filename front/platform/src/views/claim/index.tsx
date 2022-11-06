import './index.scss';
import ClaimList from './../components/ClaimList';
import Revocation from './../components/Revocation';
import { moduleActive } from '../../store/atom';
import { useRecoilState } from 'recoil';

export default function List() {

  const [activeTabStr, setActiveTabStr] = useRecoilState(moduleActive);

  return (
    <div className="template">
      {activeTabStr === 'claimList' && (
        <ClaimList />
      )}
      {activeTabStr === 'revocation' && (
        <Revocation />
      )}
    </div>
  );
}
