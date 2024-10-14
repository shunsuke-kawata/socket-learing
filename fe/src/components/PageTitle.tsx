import "../styles/components/pageTitle.css";
const PageTitle = ({ title }: { title: string }) => {
  return (
    <>
      <div className="title-div">
        <label>{title}</label>
      </div>
    </>
  );
};

export default PageTitle;
