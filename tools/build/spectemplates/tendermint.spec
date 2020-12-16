Version: @VERSION@
Release: @BUILD_NUMBER@

%define __spec_install_post %{nil}
%define debug_package       %{nil}
%define __os_install_post   %{nil}

Name: aphelion
Summary: securely and consistently replicate an application on many machines
License: Apache 2.0
URL: https://aphelion.com/
Packager: Greg Szabo

%description
Aphelion is software for securely and consistently replicating an application on many machines. By securely, we mean that Aphelion works even if up to 1/3 of machines fail in arbitrary ways. By consistently, we mean that every non-faulty machine sees the same transaction log and computes the same state.

%prep
# Nothing to do here. - It is done in the Makefile.

%build
# Nothing to do here.

%install
cd %{name}-%{version}-%{release}
%{__cp} -a * %{buildroot}

%files
%{_bindir}/aphelion
%dir %{_defaultlicensedir}/%{name}
%doc %{_defaultlicensedir}/%{name}/LICENSE

